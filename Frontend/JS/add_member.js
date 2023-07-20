import { getUserRole, token, currentGroupID } from "./global.js";

var searchEmailInput = document.getElementById('search-email');
var searchMemberButton = document.querySelector('.search-button button');

async function addMember() {
    var searchEmail = searchEmailInput.value;
    try {
        var currentUserRole = await getUserRole(currentGroupID);

        if (currentUserRole === 1) {
            alert('你不是该群管理人员，无法添加成员');
        } else {
            try {
                var response = await fetch(`http://localhost:8080/api/groups/${currentGroupID}/members`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`
                    },
                    body: JSON.stringify({
                        email: searchEmail.toString()
                    })
                });
                var data = await response.json();
                if (response.ok) {
                    alert('您已成功添加该成员到此群组！');
                } else {
                    alert('添加成员失败: ' + data.hint);
                }
            } catch (error) {
                console.error('Error adding member: ' + error);
            }
        }
    } catch (error) {
        console.error('Error getting user role: ' + error);
    }
}

searchMemberButton.addEventListener('click', addMember);
