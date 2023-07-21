import {
  UserID,
  changeTaskListHeader,
  currentGroupID, currentUserID, currentUserName,
  getUserRole,
  token,
  updateGroupMembersList, updateTaskList
} from './global.js';

(function() {
    const confirmBtn = document.querySelector('.select-team input');
    const localUser = document.querySelector('.frame-parent div .user-avatar');
    const otherUser = document.querySelector('.instance-parent');
    const selectedID = 'selected';


    confirmBtn.addEventListener('click', function(event) {
        //获取当前组的ID
    event.preventDefault();
    let select = document.getElementById('teams');
    let selectedValue = select.value;
    currentGroupID = selectedValue;
        //更新成员列表
    updateGroupMembersList(currentGroupID);
        //更新任务列表
    updateTaskList(currentGroupID, currentUserID);
     });

     otherUser.addEventListener('click', function(event) {
       if (event.target.classList.contains('user-avatar')) {
          var userInstance = event.target;
          while (userInstance && !userInstance.classList.contains('user-instance')) {
            userInstance = userInstance.parentElement;
          }
          if (userInstance) {
            var userNameElement = userInstance.querySelector('.user-name');
            currentUserID = userNameElement.value;
            var previousSelected = document.getElementById(selectedID);
            if (previousSelected) {
              previousSelected.removeAttribute('id');
            }
            userInstance.id = selectedID;
            changeTaskListHeader(e.target.parentElement.querySelector('.user-name').innerText + '的任务');
            updateTaskList(currentGroupID, currentUserID);
          }
        }
      });
      
        localUser.addEventListener('click', function(event) {
            let frameParent = event.target.parentElement;
            frameParent = frameParent.parentElement;

            let userNameElement = frameParent.querySelector('.user-name');
            currentUserID = userNameElement.value;
            var previousSelected = document.getElementById(selectedID);
            if (previousSelected) {
                previousSelected.removeAttribute('id');
            }
            frameParent.id = selectedID;
            changeTaskListHeader(currentUserName + '的任务');
            updateTaskList(currentGroupID, currentUserID);
        });
      
        otherUser.addEventListener('click', function(event) { 
          if (event.target.classList.contains('delete-user')) {
            let userRole = getUserRole(currentGroupID);
            if (userRole === 2 || userRole === 3) {
            const userID = event.target.parentElement.querySelector('.user-name').value;
            const options = {
              method: 'DELETE',
              headers: {
                'Authorization': 'Bearer ' + token
              }
            };
            fetch(`/groups/${currentGroupID}/members/${userID}`, options)
              .then(response => {
                if (!response.ok) {
                  throw new Error(`HTTP error! status: ${response.status}`);
                }
                if (currentUserID === userID) {
                  currentUserID = UserID;
                  updateTaskList(currentGroupID, currentUserID);
                }
                event.target.parentElement.parentElement.remove();
              })
              .catch(e => {
                console.log('There was a problem with your fetch operation: ' + e.message);
              });
            }
          }
        });


})();