import { token,currentGroupID, getUserRole  } from "./global";

var quitGroupButton = document.querySelector('......');
quitGroupButton.addEventListener('click',quitGroup);

async function quitGroup() {
    var currentUserRole = getUserRole(currentGroupID);

    // 如果是非群主
    if (currentUserRole === 1 || currentUserRole === 2){
        var confirmQuit = confirm('您确定要退出该群组吗？');
        if(confirmQuit) {
            try {
                var response = await fetch(`http://localhost:8080/api/groups/${currentGroupID}/members/`,{
                    method: 'DELETE',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`
                    }
                });
                if(response.ok){
                    alert('您已成功退出该群组！');
                }else {
                    var data = await response.json();
                    alert('退出群组失败： '+ data.hint);
                }
            }
            catch(error) {
                console.error('Error quitting group: '+error);
            }
        }
    }

    // 如果是群主
    else{
        var confirmDissolve = confirm('您确定要退出该群组吗？您点击确认即代表解散该群组！');
        if (confirmDissolve) {
            try {
                var response = await fetch(`http://localhost:8080/api/groups/${currentGroupID}/members/`, {
                    method: 'DELETE',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`
                    }
                });
                if (response.ok) {
                    alert('您已成功解散该群组！');
                } else {
                    var data = await response.json();
                    alert('解散群组失败： ' + data.hint);
                }
            }
            catch (error) {
                console.error('Error quitting group: ' + error);
            }
        }

    }
} 