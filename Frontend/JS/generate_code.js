import { token, getUserRole, } from "./global.js";

var createCodeButton = document.querySelector('.create-code-button');
createCodeButton.addEventListener('click',createCode);

async function createCode() {
    var groupIDInput = document.getElementById('input-group-id');
    var showCodeOutput = document.getElementById('.generate-group-code');
    var currentGroupID = groupIDInput.value;
    try {
        var currentUserRole = await getUserRole(currentGroupID);
        if(currentUserRole.ok) {
             if(currentUserRole === 1) {
        alert('你不是该群管理人员，无法生成邀请码！');
    }
        else {
        try {
            var response = await fetch(`http://localhost:8080/api/groups/${currentGroupID}/join/codes`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({
                    groupID: currentGroupID
                })
            });
            var data = await response.json();
            if (response.ok) {
                var code = data.data.code;
                showCodeOutput.value = code;
                alert('您已成功创建邀请码！请查看！');

            } else {
                alert('创建邀请码失败： ' + data.hint);
            }
        }
        catch(error) {
            console.error('Error creating code:' + error);
                }
            }
        }
    }
    catch(error) {
        console.error('Error getting role:' + error);
    }
}

