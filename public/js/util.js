window.addEventListener('load', function () {
  var content = document.getElementById("content");
  height = (Number(content.clientHeight - 50));
  var css = "margin-top:" + height + "px;";
  var footer = document.getElementById("footer").style.cssText = css;
})