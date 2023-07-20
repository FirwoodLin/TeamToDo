import {token, currentGroupID, getGroupIDs, getGroupNames, getGroupMemberNames, getGroupMemberIDs, getGroupMemberAvatars,
  getTaskNames, getTaskIDs, getTaskDescriptions, getTaskStatuses, getTaskDeadlines, getTaskStartAts, convertDateTimeFormat,
   formatDateTimeLocal, updateSelectOptions, updateGroupMembersList, updateTaskList} from "./global.js";

(function() {
    const taskModal = document.querySelector('.task-info-modal');
    const createTaskBtn = document.querySelector('.frame');
    const statusBtn = document.querySelector('.task-complete');
    const textArea = document.querySelector('.task-info-modal textarea');
    
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

})();