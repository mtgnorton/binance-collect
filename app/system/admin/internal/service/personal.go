package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

var Personal = personalService{}

type personalService struct {
	uploadAvatarPath string
}

func (p *personalService) getUploadAvatarPath() string {
	if p.uploadAvatarPath != "" {
		return p.uploadAvatarPath
	}
	p.uploadAvatarPath = g.Cfg().MustGet(gctx.New(), "upload.path").String()
	return p.uploadAvatarPath

}

func (p *personalService) Login(ctx context.Context, in define.PersonalLoginPostInput) (out *define.PersonalLoginPostOutput, err error) {
	out = &define.PersonalLoginPostOutput{}

	if false && !Common.VerifyCaptcha(ctx, in.Code, in.CaptchaId) {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "验证码错误")
		return
	}

	entity, err := Administrator.GetUserByPassportAndPassword(
		ctx,
		in.Username,
		utility.EncryptPassword(in.Username, in.Password),
	)
	if err != nil {
		return
	}

	if entity == nil {
		return out, gerror.NewCode(gcode.CodeInvalidParameter, "用户名或密码错误")
	}

	adminSummary, err := Administrator.GetAdministratorSummary(ctx, entity.Id)
	if err != nil {
		return
	}

	token, err := AdminTokenInstance.GenerateAndSaveData(ctx, in.Username, adminSummary)

	out.Token = token
	shared.Context.SetUser(ctx, adminSummary)

	return
}

func (p *personalService) Info(ctx context.Context, id uint) (out define.PersonalInfoOutput, err error) {
	out = define.PersonalInfoOutput{}
	err = dao.Administrator.Ctx(ctx).WherePri(id).Scan(&out)
	if err != nil {
		return
	}
	if out.Id == 0 {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "用户不存在")
	}
	return

}

func (p *personalService) avatarSavePath(userId uint, isAbsolute bool) string {
	uploadPath := "/upload"

	if p.getUploadAvatarPath() != "" {

		uploadPath = p.getUploadAvatarPath()
	}

	if string(uploadPath[0]) != "/" {
		uploadPath = "/" + uploadPath
	}
	if isAbsolute {
		return gfile.Join(utility.GetServerPath(), uploadPath, gconv.String(userId))
	}
	return gfile.Join(uploadPath, gconv.String(userId)) + gfile.Separator
}

func (p *personalService) Avatar(ctx context.Context, id uint, in *define.PersonAvatarInput) (out define.PersonAvatarOutput, err error) {

	out = define.PersonAvatarOutput{}
	filename, err := in.AvatarFile.Save(p.avatarSavePath(id, true), true)
	if err != nil {
		return
	}

	databasePath := p.avatarSavePath(id, false) + filename

	g.Dump(p.avatarSavePath(id, true), databasePath)
	_, err = dao.Administrator.Ctx(ctx).WherePri(id).Update(g.Map{
		dao.Administrator.Columns.Avatar: databasePath,
	})

	updatedAdministrator, err := Administrator.GetAdministratorSummary(ctx, id)

	AdminTokenInstance.UpdateData(ctx, updatedAdministrator.Username, updatedAdministrator)

	out.AvatarUrl = databasePath
	return

}

func (p *personalService) Update(ctx context.Context, administrator *model.AdministratorSummary, in *define.PersonalUpdateInput) (err error) {

	inputOldPassword := utility.EncryptPassword(administrator.Username, in.OldPassword)

	if inputOldPassword != administrator.Password {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "旧密码不正确")
		return
	}
	updateData := g.Map{
		dao.Administrator.Columns.Nickname: in.Nickname,
	}
	if in.Password != "" {
		updateData[dao.Administrator.Columns.Password] = utility.EncryptPassword(administrator.Username, in.Password)
	}
	result, err := dao.Administrator.Ctx(ctx).WherePri(administrator.Id).Fields().Update(updateData)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()

	if err != nil {
		return
	}
	if row == 0 {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "没有修改数据")
		return
	}

	updatedAdministrator, err := Administrator.GetAdministratorSummary(ctx, administrator.Id)

	AdminTokenInstance.UpdateData(ctx, updatedAdministrator.Username, updatedAdministrator)

	return
}
