
var BrowseFile = React.createClass({
  render: function() {
    return(
      <div className="browseFile">
        This is BrowseFile!
      </div>
    );
  }
});

// 上传文件参考
// http://stackoverflow.com/questions/28750489/upload-file-component-with-reactjs
var UploadFile = React.createClass({
  handleSubmit: function(e) {
    e.preventDefault();
    var formData = new FormData();
    formData.append( 'file', this.refs.file.getDOMNode().files[0] );
    var filename = React.findDOMNode(this.refs.file).value.trim();
    console.log("filename:" + filename);
    if (filename.length==0) {
      console.log("输入文件为空");
      return;
    }
    $.ajax({
      url: '/upload',  //Server script to process data
      type: 'POST',
      xhr: function() {  // Custom XMLHttpRequest
          var myXhr = $.ajaxSettings.xhr();
          if(myXhr.upload){ // Check if upload property exists
              myXhr.upload.addEventListener('progress', function(e) {
                if(e.lengthComputable){
                  console.log(e.loaded);
                  $('progress').attr({value:e.loaded,max:e.total});
              }}, false); // For handling the progress of the upload
          }
          return myXhr;
      },
      //Ajax events
      beforeSend: function () {
          console.log("beforeSend");}.bind(this),
      success: function() {
        console.log("upload success");
      }.bind(this),
      error: function(xhr, status, err) {alert("error")}.bind(this),
      // Form data
      data: formData,
      //Options to tell jQuery not to process data or worry about content-type.
      cache: false,
      contentType: false,
      processData: false
    });
  },
  render: function() {
    return (
      <div className="uploadFile">
        <form encType="multipart/form-data" onSubmit={this.handleSubmit}>
          <div className="form-group">
            <input name="file" type="file" ref="file" />
          </div>
          <div className="form-group">
            <input type="submit" value="Upload" ref="do_upload" />
          </div>
          <progress value="0"></progress>
        </form>
      </div>
    );
  }
});

var SideBar = React.createClass({
  handleChange: function(e) {
    this.setState({actived: e});
    this.props.onInputClick(e);
  },
  getInitialState: function() {
    return {
      actived: "BrowseFile"
    };
  },
  render: function() {
    var browseName = this.state.actived=="BrowseFile" ? "btn btn-primary active" : "btn btn-default";
    var uploadName = this.state.actived=="UploadFile" ? "btn btn-primary active" : "btn btn-default";
    return (
      <div className="sideBar col-xs-3 col-md-1">
        <input className={browseName} type="button" value="文件浏览" onClick={this.handleChange.bind(this, "BrowseFile")} /><br/>
        <input className={uploadName} type="button" value="文件上传" onClick={this.handleChange.bind(this, "UploadFile")} />
      </div>
    );
  }
});

var MenuResponse = React.createClass({
  render: function() {
    var key = this.props.inputName;
    var response = RESPONSES[key];
    return (
      <div className="menuResponse col-xs-9 col-md-6">
        {response}
      </div>
    );
  }
});

var NetFile = React.createClass({
  handleInputClick: function(e) {
    console.log("set state " + e);
    this.setState({inputName: e});
  },
  getInitialState: function() {
    return {
      inputName: "BrowseFile"
    };
  },
  render: function() {
    return (
      <div className="netFile row">
        <SideBar onInputClick={this.handleInputClick} />
        <MenuResponse inputName={this.state.inputName} />
      </div>
    );
  }
});

var RESPONSES = {BrowseFile: <BrowseFile />, UploadFile: <UploadFile />};

React.render(
  <NetFile />,
  document.getElementById('content')
);
