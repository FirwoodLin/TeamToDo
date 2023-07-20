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

function getTaskStatuses(groupID, userID) {
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
      var taskStatuses = data.data.tasks.map(task => task.taskStatus);
      return taskStatuses;
    } else {
      console.log("请求未成功: " + data.hint);
    }
  })
  .catch(e => {
    console.log("请求出错: " + e);
  });
};

function getTaskDeadlines(groupID, userID) {
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
      var taskDeadlines = data.data.tasks.map(task => task.deadline);
      taskDeadlines = taskDeadlines.map(deadline => convertDateTimeFormat(deadline));
      return taskDeadlines;
    } else {
      console.log("请求未成功: " + data.hint);
    }
  })
  .catch(e => {
    console.log("请求出错: " + e);
  });
};

function getTaskStartAts(groupID, userID) {
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
       var taskStartAts = data.data.tasks.map(task => task.startAt);
       taskStartAts = taskStartAts.map(startAt => convertDateTimeFormat(startAt));
       return taskDeadlines;
     } else {
       console.log("请求未成功: " + data.hint);
     }
   })
   .catch(e => {
     console.log("请求出错: " + e);
   });
 };

function convertDateTimeFormat(inputDateTime) {
  var date = new Date(inputDateTime);
  var formattedDate = date.toLocaleString('en-GB', {
      hour: '2-digit',
      minute: '2-digit',
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
  });
  var parts = formattedDate.split(',');
  var outputDateTime = parts[1].trim() + ', ' + parts[0].trim();

  return outputDateTime;
}

function formatDateTimeLocal(inputDateTimeLocal) {
  // Create a Date object from the input string
  var date = new Date(inputDateTimeLocal);

  // Format the date and time parts
  var year = date.getFullYear();
  var month = (date.getMonth() + 1).toString().padStart(2, '0');
  var day = date.getDate().toString().padStart(2, '0');
  var hours = date.getHours().toString().padStart(2, '0');
  var minutes = date.getMinutes().toString().padStart(2, '0');
  var seconds = date.getSeconds().toString().padStart(2, '0');
  var milliseconds = date.getMilliseconds().toString().padStart(3, '0');

  // Format the timezone offset
  var offset = -date.getTimezoneOffset();
  var offsetSign = offset >= 0 ? '+' : '-';
  offset = Math.abs(offset);
  var offsetHours = Math.floor(offset / 60).toString().padStart(2, '0');
  var offsetMinutes = (offset % 60).toString().padStart(2, '0');

  // Join the parts into an ISO 8601 string
  var outputDateTime = `${year}-${month}-${day}T${hours}:${minutes}:${seconds}.${milliseconds}${offsetSign}${offsetHours}:${offsetMinutes}`;

  return outputDateTime;
}

//export
export { token, getGroupIDs, getGroupNames, getGroupMemberNames, getGroupMemberIDs, getGroupMemberAvatars,
  getTaskNames, getTaskIDs, getTaskDescriptions, getTaskStatuses, getTaskDeadlines, convertDateTimeFormat, formatDateTimeLocal};

