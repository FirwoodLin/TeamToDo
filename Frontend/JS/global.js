//global variable
var token;


//global function
function getGroupIDs() {
    fetch('/groups', {
      method: 'GET',
      headers: {
        'Authorization': 'Bearer ' + token,
      },
    })
    .then(response => {
      if (!response.ok) {
        throw new Error("HTTP error " + response.status);
      }
      return response.json();
    })
    .then(json => {
      if (json.success) {
        let groupIDs = json.data.groups.map(group => group.groupID);
        return groupIDs;
      } else {
        throw new Error(json.hint); 
      }
    })
    .catch(error => {
      console.error('An error occurred:', error);
    });
};

function getGroupNames() {
    fetch('/groups', {
        method: 'GET',
        headers: {
          'Authorization': 'Bearer ' + token,
        },
      })
      .then(response => {
        if (!response.ok) {
          throw new Error("HTTP error " + response.status);
        }
        return response.json();
      })
      .then(json => {
        if (json.success) {
          let groupNames = json.data.groups.map(group => group.groupName);
          return groupNames;
        } else {
          throw new Error(json.hint); 
        }
      })
      .catch(error => {
        console.error('An error occurred:', error);
      });
};

function getGroupMemberNames(groupID) {
  fetch(`/groups/${groupID}/members`, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`
    }
  })
  .then(response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
  })
  .then(data => {
    if (data.success) {
      var userNames = data.data.map(item => item.User.userName);
      return userNames;
    } else {
      console.log("请求未成功: " + data.hint);
    }
  })
  .catch(e => {
    console.log("请求出错: " + e);
  });
};

function getGroupMemberIDs(groupID) {
  fetch(`/groups/${groupID}/members`, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`
    }
  })
  .then(response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
  })
  .then(data => {
    if (data.success) {
      var userIDs = data.data.map(item => item.User.userID);
      return userIDs;
    } else {
      console.log("请求未成功: " + data.hint);
    }
  })
  .catch(e => {
    console.log("请求出错: " + e);
  });
};

function getGroupMemberAvatars(groupID) {
  fetch(`/groups/${groupID}/members`, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`
    }
  })
  .then(response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
  })
  .then(data => {
    if (data.success) {
      var userAvatars = data.data.map(item => item.User.userAvatar);
      return userIDs;
    } else {
      console.log("请求未成功: " + data.hint);
    }
  })
  .catch(e => {
    console.log("请求出错: " + e);
  });
};

//可以重载,依据传递的参数数量不同,执行不同的操作,后面的其他信息也是同理
function getTaskNames(groupID, userID) {
  fetch(`/tasks?groupID=${groupID}&userID=${userID}`, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`
    }
  })
  .then(response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
  })
  .then(data => {
    if (data.success) {
      var taskNames = data.data.tasks.map(task => task.taskName);
      return taskNames;
    } else {
      console.log("请求未成功: " + data.hint);
    }
  })
  .catch(e => {
    console.log("请求出错: " + e);
  });
};

function getTaskIDs(groupID, userID) {
  fetch(`/tasks?groupID=${groupID}&userID=${userID}`, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`
    }
  })
  .then(response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
  })
  .then(data => {
    if (data.success) {
      var taskIDs = data.data.tasks.map(task => task.taskID);
      return taskIDs;
    } else {
      console.log("请求未成功: " + data.hint);
    }
  })
  .catch(e => {
    console.log("请求出错: " + e);
  });
};

function getTaskDescriptions(groupID, userID) {
  fetch(`/tasks?groupID=${groupID}&userID=${userID}`, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`
    }
  })
  .then(response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
  })
  .then(data => {
    if (data.success) {
      var taskDescriptions = data.data.tasks.map(task => task.description);
      return taskDescriptions;
    } else {
      console.log("请求未成功: " + data.hint);
    }
  })
  .catch(e => {
    console.log("请求出错: " + e);
  });
};





//export
export { token, getGroupIDs, getGroupNames, getGroupMemberNames, getGroupMemberIDs, getGroupMemberAvatars,
  getTaskNames, getTaskIDs, getTaskDescriptions};

