import { token } from "./global.js";


const joinTeamModal = document.getElementById('joinTeamModal');
const createTeamModal = document.getElementById('createTeamModal');

var joinbtn = document.querySelector('.join-team-btn');
var createbtn = document.querySelector('.create-team-btn');

// 默认隐藏
document.addEventListener('DOMContentLoaded', hideAll);
function hideAll() {
    joinTeamModal.style.display = 'none';
    createTeamModal.style.display = 'none';
};

// 显示加入团队
function showJoinTeamModal() {
    joinTeamModal.style.display = 'block';
    createTeamModal.style.display = 'none';
}
joinbtn.addEventListener('click',showJoinTeamModal);

// 显示创建团队
function showCreateTeamModal() {
    joinTeamModal.style.display = 'none'; 
    createTeamModal.style.display = 'block';
}
createbtn.addEventListener('click',showCreateTeamModal);

// 退出modal按钮响应
var quitbtn1 = document.querySelector('.quit-btn1');
var quitbtn2 = document.querySelector('.quit-btn2');
quitbtn1.addEventListener('click',hideAll);
quitbtn2.addEventListener('click',hideAll);


// 创建团队逻辑
var groupNameInput = document.getElementById('group-name');
var groupDescription = document.getElementById('group-description');
var groupIDInput = document.getElementById('group-ID');

async function CreateGroup() {
    var groupName = groupNameInput.value;
    var description = groupDescription.value;

    try {
        var response = await fetch('http://localhost:8080/api/groups', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization':`Bearer ${token}`
            },
            body: JSON.stringify({
                groupName: groupName,
                description: description
            })
        });

        var data = await response.json();
        if(response.ok) {
            var groupID = data.data.groupID;
            groupIDInput.value = groupID;
        }else {
            alert('啊哦~出现了一些错误');
        }
    }
    catch (error) {
        console.error('Error creating group:', error);
        alert('出错啦 请稍后再试');
    }
}

var generateGroupButton = document.querySelector('.search-button2 button');
generateGroupButton.addEventListener('click', createGroup); 
