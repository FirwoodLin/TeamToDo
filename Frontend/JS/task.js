import {token, currentGroupID, currentUserID, currentUserName, currentUserEmail, currentUserAvatar,getGroupIDs, getGroupNames, getGroupMemberNames, getGroupMemberIDs, getGroupMemberAvatars,
  getTaskNames, getTaskIDs, getTaskDescriptions, getTaskStatuses, getTaskDeadlines, getTaskStartAts, convertDateTimeFormat,
  formatDateTimeLocal, updateSelectOptions, updateGroupMembersList, updateTaskList, getUserRole, formatDateTimeLocalToClient, convertClientTimeToDateLocal
} from "./global.js";

(function() {
    const taskModal = document.querySelector('.task-info-modal');
    const statusBtn = document.querySelector('.task-complete');
    const textArea = document.querySelector('.task-info-modal textarea');
    const saveBtn = document.querySelector('.save-task');
    var isUpdating = false;
    
    //文本编辑模块框打开
    $(function() {
        $('#open-task-modal').click(function(event) {
          event.preventDefault();
          $('#task-info-modal').modal({
            modalClass: 'task-info-modal'
          });

          let isCompleted = document.querySelector('.task-complete');
          isCompleted.dataset.state = '0';
        });
      });
    
      $(function() {
        $('.task-list').on('dblclick', '.task', function(event) {
            event.preventDefault();
            // 打开模态窗口
            $('#task-info-modal').modal({
                modalClass: 'task-info-modal'
            });
              //更新完成按钮的状态
          let isCompleted = document.querySelector('.task-complete');
          var taskStatus = $(this).find('.item4').attr('data-state');
          isCompleted.dataset.state = taskStatus;
            // 获取任务信息
        var taskName = $(this).find('.item1').text();
        var taskStartTime = $(this).find('.item2').text();
        var taskEndTime = $(this).find('.item3').text();
        var taskDescription = $(this).find('.item5').attr('value');

        taskStartTime = convertClientTimeToDateLocal(taskStartTime);
        taskEndTime = convertClientTimeToDateLocal(taskEndTime);
        // 将任务信息填入模态窗口的对应输入框中
        $('#task-name').val(taskName);
        $('#task-start-time').val(taskStartTime);
        $('#task-end-time').val(taskEndTime);
        $('#task-description').val(taskDescription);
        // 保存按钮变为更新按钮
        isUpdating = true;
        });
    });


    //文本编辑模块框关闭
    $(function() {
        $('.save-task').click(function(event) {
          event.preventDefault();
          $.modal.close();
        });
      });

      //自动调整文本编辑框大小
    textArea.addEventListener('input', function() {
        this.style.height = 'auto';
        this.style.height = (this.scrollHeight) + 'px';
    });

    //保存任务,追加到任务列表
    saveBtn.addEventListener('click', function(event) {
        event.preventDefault();
        if(isUpdating) {
        updateTask();
        }
        else {
          addTask();
      }
        isUpdating = false;
    });
    
    //完成按钮
    statusBtn.addEventListener('click', function(event) {
        event.preventDefault();
        if (this.dataset.state === '0') {
          this.dataset.state = '1';
        } else {
          this.dataset.state = '0';
        }
      });

      async function addTask() {
        await sendTaskToServer(); 
        clearTaskInput();
        await updateTaskList(currentGroupID, currentUserID);
      }

      async function updateTask() {
        await updateTaskToServer();
        clearTaskInput();
        await updateTaskList(currentGroupID, currentUserID);
      }
    // function createTask(taskName, taskStartAt, taskDeadline, taskDescription) {
    //   let taskList = document.querySelector('.task-list');
    //   taskStartAt = formatDateTimeLocalToClient(taskStartAt);
    //   taskDeadline = formatDateTimeLocalToClient(taskDeadline);

    //   let li = document.createElement('li');
    //   li.className = 'task';
  
    //   let div = document.createElement('div');
    //   div.className = 'delete-task';
    //   li.appendChild(div);
  
    //   let p1 = document.createElement('p');
    //   p1.className = 'item1';
    //   p1.textContent = taskName;
    //   p1.value = null;
    //   li.appendChild(p1);
  
    //   let p2 = document.createElement('p');
    //   p2.className = 'item2';
    //   p2.textContent = taskStartAt;
    //   li.appendChild(p2);
  
    //   let p3 = document.createElement('p');
    //   p3.className = 'item3';
    //   p3.textContent = taskDeadline;
    //   li.appendChild(p3);
  
    //   let p4 = document.createElement('p');
    //   p4.className = 'item4';
    //   p4.textContent = '未完成';
    //   p4.value = 0;
    //   li.appendChild(p4);

    //   let p5 = document.createElement('p');
    //   p5.className = 'item5';
    //   p5.value = taskDescription;
  
    //   taskList.appendChild(li);
    // }

    function clearTaskInput() {
      document.getElementById('task-name').value = '';
      document.getElementById('task-start-time').value = '';
      document.getElementById('task-end-time').value = '';
      document.getElementById('task-description').value = '';
  }

  function getUserInput() {
    var taskName = document.getElementById('task-name').value;
    var taskStartTime = document.getElementById('task-start-time').value;
    var taskEndTime = document.getElementById('task-end-time').value;
    var taskDescription = document.getElementById('task-description').value;

    return [taskName, taskStartTime, taskEndTime, taskDescription];
}

async function sendTaskToServer() {
  var [taskName, startAt, deadline, description] = getUserInput();

  var body = {
      taskName: taskName,
      description: description,
      startAt: startAt,
      deadline: deadline,
      taskStatus: 0, 
      groupID: currentGroupID 
  };

  try {
      var response = await fetch("/tasks", {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
              'Authorization': 'Bearer ' + token 
          },
          body: JSON.stringify(body)
      });

      if (!response.ok) {
          throw new Error("HTTP error " + response.status);
      }
      var data = await response.json();
      return data;
  } catch (error) {
      console.error("Failed to send task to server: ", error);
  }
}

async function updateTaskToServer() {

  var [taskName, startAt, deadline, description] = getUserInput();
  var taskStatus = +(document.querySelector('.task-complete').dataset.state);

  startAt = formatDateTimeLocal(startAt);
  deadline = formatDateTimeLocal(deadline);

  try {
    const response = await fetch(`/tasks/${taskID}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token
        },
        body: JSON.stringify({
            taskName: taskName,
            description: description,
            startAt: startAt,
            deadline: deadline,
            taskStatus: taskStatus
        })
    });

    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    return data;
} catch (error) {
    console.error('There was a problem with the fetch operation:', error);
}
}

})();