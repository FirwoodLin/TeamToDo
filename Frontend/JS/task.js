import {
  formatDateTimeLocalToClient
} from "./global.js";

(function() {
    const taskModal = document.querySelector('.task-info-modal');
    const createTaskBtn = document.querySelector('.frame');
    const statusBtn = document.querySelector('.task-complete');
    const textArea = document.querySelector('.task-info-modal textarea');
    const saveBtn = document.querySelector('.save-task');
    
    //文本编辑模块框打开
    $(function() {
        $('#open-task-modal').click(function(event) {
          event.preventDefault();
          $('#task-info-modal').modal({
            modalClass: 'task-info-modal'
          });
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
    });
    

    function createTask(taskName, taskStartAt, taskDeadline) {
      let taskList = document.querySelector('.task-list');
      taskStartAt = formatDateTimeLocalToClient(taskStartAt);
      taskDeadline = formatDateTimeLocalToClient(taskDeadline);

      let li = document.createElement('li');
      li.className = 'task';
  
      let div = document.createElement('div');
      div.className = 'delete-task';
      li.appendChild(div);
  
      let p1 = document.createElement('p');
      p1.className = 'item1';
      p1.textContent = taskName;
      li.appendChild(p1);
  
      let p2 = document.createElement('p');
      p2.className = 'item2';
      p2.textContent = taskStartAt;
      li.appendChild(p2);
  
      let p3 = document.createElement('p');
      p3.className = 'item3';
      p3.textContent = taskDeadline;
      li.appendChild(p3);
  
      let p4 = document.createElement('p');
      p4.className = 'item4';
      p4.textContent = '未完成';
      p4.value = 0;
      li.appendChild(p4);
  
      taskList.appendChild(li);
    }
})();