//global variable
var token;
var currentGroupID = initCurrentGroupID();
var currentUserID, currentUserName, currentUserEmail, currentUserAvatar;
var UserID;


//global function

// 查询所有群组ID
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

// 查询所有群组名
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

// 查询群组所有成员名
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

// 查询群组所有成员ID
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

// 查询群组所有成员Avatar
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

  var outputDateTime = `${year}-${month}-${day}T${hours}:${minutes}:${seconds}.${milliseconds}${offsetSign}${offsetHours}:${offsetMinutes}`;

  return outputDateTime;
}

function updateSelectOptions() {
  let groupNames = getGroupNames();
  let groupIDs = getGroupIDs();
  let select = document.getElementById('teams');

  select.innerHTML = '';

  for (let i = 0; i < groupNames.length; i++) {
      let option = document.createElement('option');
      option.value = groupIDs[i];
      option.text = groupNames[i];
      select.appendChild(option);
  }
}



function updateGroupMembersList(groupID) {
  // 调用相应的函数获取新的成员名称，ID 和头像
  let memberNames = getGroupMemberNames(groupID);
  let memberIDs = getGroupMemberIDs(groupID);
  let memberAvatars = getGroupMemberAvatars(groupID);

  let list = document.querySelector('.instance-parent');

  list.innerHTML = '';

  for (let i = 0; i < memberNames.length; i++) {
      let listItem = document.createElement('li');
      listItem.className = 'user-instance';

      let userInfo = document.createElement('div');
      userInfo.className = 'user-info';

      let avatar = document.createElement('img');
      avatar.className = 'user-avatar';
      avatar.src = memberAvatars[i];
      avatar.alt = '用户头像';

      let name = document.createElement('div');
      name.className = 'user-name';
      name.textContent = memberNames[i];

      name.value = memberIDs[i];

      let deleteUser = document.createElement('div');
      deleteUser.className = 'delete-user';

      userInfo.appendChild(avatar);
      userInfo.appendChild(name);
      userInfo.appendChild(deleteUser);

      listItem.appendChild(userInfo);

      list.appendChild(listItem);
  }
}

function updateTaskList(groupID, userID) {
  let taskNames = getTaskNames(groupID, userID);
  let taskIDs = getTaskIDs(groupID, userID);
  let taskStatuses = getTaskStatuses(groupID, userID);
  let taskDeadlines = getTaskDeadlines(groupID, userID);
  let taskStartAts = getTaskStartAts(groupID, userID);

  let taskList = document.querySelector('.task-list');

  while (taskList.firstChild) {
    taskList.removeChild(taskList.firstChild);
  }

  for (let i = 0; i < taskNames.length; i++) {
    let li = document.createElement('li');
    li.className = 'task';
    li.id = i;

    let div = document.createElement('div');
    div.className = 'delete-task';
    li.appendChild(div);

    let p1 = document.createElement('p');
    p1.className = 'item1';
    p1.textContent = taskNames[i];
    p1.value = taskIDs[i];  // Set taskID as value
    li.appendChild(p1);

    let p2 = document.createElement('p');
    p2.className = 'item2';
    p2.textContent = taskStartAts[i];
    li.appendChild(p2);

    let p3 = document.createElement('p');
    p3.className = 'item3';
    p3.textContent = taskDeadlines[i];
    li.appendChild(p3);

    let p4 = document.createElement('p');
    p4.className = 'item4';
    p4.textContent = taskStatuses[i];
    li.appendChild(p4);

    taskList.appendChild(li);
  }
}

function initCurrentGroupID() {
  let firstGroupID = getGroupIDs()[0];
  return firstGroupID;
}

// 查询成员在某群组的身份
async function getUserRole(groupID) {
  try {
    var response = await fetch(`http://localhost:8080/api/groups/${groupID}`,{
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      }
    });
    var data = await response.json();
    if(response.ok) {
      var userRole = data.data.role;
      console.log('用户的身份是: '+userRole);
      return userRole;
    }else {
      alert('查询失败: '+ data.hint);
    }
  }
  catch(error) {
    console.error('Error querying user role:'+ error);
  }
}

//export
export {
  token, currentGroupID, currentUserID, currentUserName, currentUserEmail, currentUserAvatar,UserID,getGroupIDs, getGroupNames, getGroupMemberNames, getGroupMemberIDs, getGroupMemberAvatars,
  getTaskNames, getTaskIDs, getTaskDescriptions, getTaskStatuses, getTaskDeadlines, getTaskStartAts, convertDateTimeFormat,
   formatDateTimeLocal, updateSelectOptions, updateGroupMembersList, updateTaskList, getUserRole};

