var NetFile = React.createClass({
  render: function() {
    return (
      <div className="netFile row">
        <SideBar />
        <MenuResponse />
      </div>
    );
  }
});

var SideBar = React.createClass({
  browseFile: function() {
    console.log("browseFile call");
  },
  uploadFile: function() {
    console.log("uploadFile call")
  },
  render: function() {
    return (
      <div className="sideBar col-xs-3 col-md-1">
        <input className="btn btn-default" type="button" value="文件浏览" onClick={this.browseFile} /><br/>
        <input className="btn btn-default" type="button" value="文件上传" onClick={this.uploadFile}/>
      </div>
    );
  }
});

var MenuResponse = React.createClass({
  render: function() {
    return (
      <div className="menuResponse col-xs-9 col-md-6">
        MenuResponse!!
      </div>
    );
  }
});

React.render(
  <NetFile />,
  document.getElementById('content')
);
