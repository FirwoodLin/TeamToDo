import { token, currentGroupID, getGroupIDs, getGroupNames, getGroupMemberNames, getGroupMemberIDs, getGroupMemberAvatars,
    getTaskNames, getTaskIDs, getTaskDescriptions, getTaskStatuses, getTaskDeadlines, getTaskStartAts, convertDateTimeFormat,
     formatDateTimeLocal, updateSelectOptions, updateGroupMembersList, updateTaskList} from "./global.js";

(function() {
    const confirmBtn = document.querySelector('.select-team input');



    confirmBtn.addEventListener('click', function(event) {
        //获取当前组的ID
    event.preventDefault();
    let select = document.getElementById('teams');
    let selectedValue = select.value;
    currentGroupID = selectedValue;
        //获取当前组的成员ID
    let groupMemberIDs = getGroupMemberIDs(selectedValue);
        //更新成员列表
    updateGroupMembersList(currentGroupID);
        //更新任务列表
    updateTaskList(currentGroupID, );
     });





})();