import {
  token, currentGroupID, currentUserID, currentUserName, currentUserEmail, currentUserAvatar, UserID, getGroupIDs, getGroupNames, getGroupMemberNames, getGroupMemberIDs, getGroupMemberAvatars,
  getTaskNames, getTaskIDs, getTaskDescriptions, getTaskStatuses, getTaskDeadlines, getTaskStartAts, convertDateTimeFormat,
  formatDateTimeLocal, updateSelectOptions, updateGroupMembersList, updateTaskList, getUserRole, formatDateTimeLocalToClient, convertClientTimeToDateLocal
} from "./global.js";

(function () {
  const confirmBtn = document.querySelector('.select-team input');
  const localUser = document.querySelector('.frame-parent div .user-avatar');
  const otherUser = document.querySelector('.instance-parent');
  const selectedID = 'selected';


  confirmBtn.addEventListener('click', function (event) {
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

  otherUser.addEventListener('click', function (event) {
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
        updateTaskList(currentGroupID, currentUserID);
      }
    }
  });

  localUser.addEventListener('click', function (event) {
    let frameParent = event.target.parentElement;
    frameParent = frameParent.parentElement;

    let userNameElement = frameParent.querySelector('.user-name');
    currentUserID = userNameElement.value;
    var previousSelected = document.getElementById(selectedID);
    if (previousSelected) {
      previousSelected.removeAttribute('id');
    }
    frameParent.id = selectedID;
    updateTaskList(currentGroupID, currentUserID);
  });


})();


var generateCodeModal = document.querySelector('.generate-code-box')

// 默认隐藏模态窗口
document.addEventListener('DOMContentLoaded', hideAll);
function hideAll() {
  console.log('hideAll 函数被调用');
  generateCodeModal.style.display = 'none';
};

var generateCodeBtn = document.querySelector('.create-code-btn')
// 点击按钮显示模态窗口
function showGenerateCodeModal() {
  generateCodeModal.style.display = 'block';
}
generateCodeBtn.addEventListener('click', showGenerateCodeModal);

/* 退出模态窗口 */
var quitbtn = document.querySelector('.quit-btn');
quitbtn.addEventListener('click', hideAll);
console.log("yes");

// 呈现个人信息窗口// 呈现个人信息窗口
var groupIcon = document.querySelector('.group-icon');
var personalBox = document.querySelector('.personal-box');
var timer; 

groupIcon.addEventListener('mouseover', function () {
  clearTimeout(timer); 
  personalBox.style.display = 'block';

  personalBox.addEventListener('mouseout', function (event) {
    if (!personalBox.contains(event.toElement)) {
      personalBox.style.display = 'none';
    }
  });
});

personalBox.addEventListener('mouseover', function () {
  clearTimeout(timer); 
});

groupIcon.addEventListener('mouseout', function () {
  timer = setTimeout(function () {
    if (!personalBox.contains(document.querySelector(':hover'))) {
      personalBox.style.display = 'none';
    }
  }, 1000); 
});