

var getUrlParameter = function getUrlParameter(sParam) {
    var sPageURL = window.location.search.substring(1),
        sURLVariables = sPageURL.split('&'),
        sParameterName,
        i;

    for (i = 0; i < sURLVariables.length; i++) {
        sParameterName = sURLVariables[i].split('=');

        if (sParameterName[0] === sParam) {
            return sParameterName[1] === undefined ? true : decodeURIComponent(sParameterName[1]);
        }
    }
    return false;
};


function notSupport(){
    alert("该功能暂未开放！")
}

$(document).ready(function() {

    // Check for click events on the navbar burger icon
    // $(".navbar-burger").click(function() {
    //
    //     $(".navbar-burger").toggleClass("is-active");
    //     $(".navbar-menu").toggleClass("is-active");
    //
    // });
});
