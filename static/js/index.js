
var BrowseFile = React.createClass({
  render: function() {
    return(
      <div className="browseFile">
        This is BrowseFile!
      </div>
    );
  }
});

var UploadFile = React.createClass({
  render:function() {
    return(
      <div className="uploadFile">
        This is UploadFile!
      </div>
    );
  }
});

var SideBar = React.createClass({
  handleChange: function(e) {
    this.props.onInputClick(e);
  },
  render: function() {
    return (
      <div className="sideBar col-xs-3 col-md-1">
        <input className="btn btn-default" type="button" value="文件浏览" onClick={this.handleChange.bind(this, "BrowseFile")} /><br/>
        <input className="btn btn-default" type="button" value="文件上传" onClick={this.handleChange.bind(this, "UploadFile")}/>
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
