function progressHandlingFunction(e){
    if(e.lengthComputable){
        $('progress').attr({value:e.loaded,max:e.total});
    }
  }
$(document).ready(function(){
  $("progress").hide();
  $(".alert").hide();
  $('#file').change(function(){
    var file = this.files[0];
    var name = file.name;
    var size = file.size;
    var type = file.type;
    $("#upload_alert_warn").hide();
    $("#upload_alert_ok").hide();
    //Your validation
  });
  $('#upload').click(function(){
    var formData = new FormData($('form')[0]);
    var filename = $("#file").val();
    if (filename.length==0) {
      console.log("输入文件为空");
      $("#upload_alert_warn").show();
      return;
    }
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
        beforeSend: function () {
            console.log("beforeSend");
            $("progress").show();},
        success: function() {
          console.log("upload success");
          $("#upload_alert_ok").show();
        },
        error: function() {$("progress").hide();alert("error")},
        // Form data
        data: formData,
        //Options to tell jQuery not to process data or worry about content-type.
        cache: false,
        contentType: false,
        processData: false
    });
  });
  $('#upload_alert_warn_close').click(function() {
    $('#upload_alert_warn').hide();
  });
  $('#upload_alert_ok_close').click(function() {
    $('#upload_alert_ok').hide();
  });
});

// HTML元素的名称以小写字母开头，React class的名称以大写字母开头
var FileinfoBox = React.createClass({
  getInitialState: function() {
    return {data: []};
  },
  componentDidMount: function() {
    $.ajax({
      url: this.props.url,
      type: 'GET',
      data: {browsePath: this.props.browsePath},
      dataType: 'json',
      cache: false,
      success: function(data) {
        console.log(data)
        this.setState({data: data});
      }.bind(this),
      error: function(xhr, status, err) {
        console.error(this.props.url, status, err.toString());
      }.bind(this)
    });
  },
  render: function() {
    return (
      <div className="fileinfoBox">
        <FileinfoList data={this.state.data} />
      </div>
    );
  }
});

var FileinfoList = React.createClass({
  render: function() {
    var fileNodes = this.props.data.map(function(fileinfo) {
      return (
        <FileinfoItem name={fileinfo.name} path={fileinfo.path}></FileinfoItem>
      );
    });
    return (
      <div className="fileinfoList">
        <ul>
          {fileNodes}
        </ul>
      </div>
    );
  }
});


var FileinfoItem = React.createClass({
  render: function() {
    var urlpath = "/download?file=" + this.props.path;
    return (
      <div className="fileinfoItem">
        <li><a href={urlpath}> {this.props.name} </a></li>
      </div>
    );
  }
});

React.render(
  <FileinfoBox url="getfiles" browsePath="./files"/>,
  document.getElementById('fileinfos')
);
