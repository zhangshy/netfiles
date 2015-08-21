function progressHandlingFunction(e){
    if(e.lengthComputable){
        $('progress').attr({value:e.loaded,max:e.total});
    }
  }
$(document).ready(function(){
  $('#upload').click(function(){
    var formData = new FormData($('form')[0]);
    $.ajax({
        url: 'upload',  //Server script to process data
        type: 'POST',
        xhr: function() {  // Custom XMLHttpRequest
            var myXhr = $.ajaxSettings.xhr();
            if(myXhr.upload){ // Check if upload property exists
                myXhr.upload.addEventListener('progress',progressHandlingFunction, false); // For handling the progress of the upload
            }
            return myXhr;
        },
        //Ajax events
        //beforeSend: function () {alert("beforeSend")},
        //success: function() {alert("success")},
        error: function() {alert("error")},
        // Form data
        data: formData,
        //Options to tell jQuery not to process data or worry about content-type.
        cache: false,
        contentType: false,
        processData: false
    });
  });
  $.ajax({
    url: 'getfiles',
    type: 'GET',
    data: {browsePath:'D:\\BaiduYunDownload'},
    success: function(data) {
      $("#filelists").html(data);
    },
    error: function(xhr, status, err) {
      alert("error!")
    }
  });
});