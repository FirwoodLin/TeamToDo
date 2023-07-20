import { getUserRole, token, currentGroupID } from "./global.js";

var searchEmailInput = document.getElementById('search-email');
var searchMemberButton = document.querySelector('.search-button button');

function addMember() {
    var searchEmail = searchEmailInput.value;
    var currentGroupID = currentGroupID


}
searchMemberButton.addEventListener('click',addMember);