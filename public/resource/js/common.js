var gf = {
  setError: function (message) {
    $(".problem").show();
    $(".problem ul").empty();
    $(".problem ul").append("<li>" + message + "</li>");
  },

  setCookie: function setCookie(name, value, hours) {
    var d = new Date();
    d.setTime(d.getTime() + (hours * 60 * 60 * 1000));
    var expires = "expires=" + d.toGMTString();
    document.cookie = name + "=" + value + "; " + expires;
  }
}
