import { token } from "./global.js";


const joinTeamModal = document.getElementById('joinTeamModal');
const createTeamModal = document.getElementById('createTeamModal');

var joinbtn = document.querySelector('.join-team-btn');
var createbtn = document.querySelector('.create-team-btn');

// 个人主页显示信息
function showInformation() {
    var userName = localStorage.getItem('userName');
    var userEmail = localStorage.getItem('userEmail');

    var userNameElement = document.getElementById('user-name');
    var userEmailElement = document.getElementById('user-email');

    userNameElement.textContent = userName;
    userEmailElement.textContent = userEmail;
}
document.addEventListener('DOMContentLoaded',showInformation);

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
            alert('您已成功创建新群组！');
            hideAll();
        }else {
            alert('啊哦~出现了一些错误');
        }
    }
    catch (error) {
        console.error('Error creating group:', error);
    }
}

var generateGroupButton = document.querySelector('.search-button2 button');
generateGroupButton.addEventListener('click', CreateGroup); 


// 通过群组ID加入群组
var searchGroupID = document.getElementById('search-groupID');
async function joinGroup() {
    var groupID = searchGroupID.value;

    try {
        var response = await fetch('http://localhost:8080/api/groups/join', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
            body: JSON.stringify({
                groupID: groupID
            })
        });
        if(response.ok){
            alert('您已成功加入该群组！');
            hideAll();
        }else {
            alert('加入群组失败！请检查您的信息是否有误！')
        }
    }
    catch (error) {
        console.error('Error joining group:', error);

    }

}
var joinTeamButton = document.querySelector('.search-button1 button');
joinTeamButton.addEventListener('click',joinGroup);