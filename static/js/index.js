
var BrowseFile = React.createClass({
  handleSubmit: function(e) {
    e.preventDefault();
    root = this.refs.pathName.getDOMNode().value;
    console.log(root);
    if (root.substring(0, 2).toLowerCase()=="C:".toLowerCase()) {
      alert("不允许浏览C盘");
      return;
    }
    // 使用jquery参考
    // http://stackoverflow.com/questions/25436445/using-jquery-plugins-that-transform-the-dom-in-react-components
    var $tree = $(this.refs.tree_using_ajax.getDOMNode());
    $tree.jstree("refresh"); // 刷新jstree
    $tree.jstree({
      'core' : {
        'data' : {
          'url' : "/tree_file",
          "dataType" : "json",
          'data' : function (node) {
            console.log("using_ajax node id: " + node.id + " root:" + root)
            if (node.id==='#') {
              return { 'id' : root };
            } else {
              return { 'id' : node.id };
            }
          }
        }
      }
    });
  },
  render: function() {
    return(
      <div className="browseFile">
        <form onSubmit={this.handleSubmit} >
          <input type="text" name="pathName" placeholder="请输入要浏览的文件路径" required="required" ref="pathName" />
          <input type="submit" value="确定" />
        </form>
        <div ref="tree_using_ajax" ></div>
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
            <input name="file" type="file" required="required" ref="file" />
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
  render: function() {
    var browseName = this.props.actived=="BrowseFile" ? "active" : "";
    var uploadName = this.props.actived=="UploadFile" ? "active" : "";
    return (
      <div className="sideBar col-xs-3 col-md-1">
        <ul className="nav nav-pills nav-stacked">
          <li role="presentation" className={browseName} onClick={this.handleChange.bind(this, "BrowseFile")} ><a href="#">文件浏览</a></li>
          <li role="presentation" className={uploadName} onClick={this.handleChange.bind(this, "UploadFile")} ><a href="#">文件上传</a></li>
        </ul>
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
        <SideBar onInputClick={this.handleInputClick} actived={this.state.inputName} />
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
