(function() {
    var taskModal = document.querySelector('.task-info-modal');
    var createTaskBtn = document.querySelector('.frame');
    var closeBtn = document.querySelector('.task-complete');
    var textArea = document.querySelector('.task-info-modal textarea');
    
    createTaskBtn.addEventListener('click', function() {
        taskModal.style.display = 'block';
        console.log('成功打开模态窗口');
    });

    closeBtn.addEventListener('click', function() {
        taskModal.style.display = 'none';
        console.log('成功关闭模态窗口');
    });

    textArea.addEventListener('input', function() {
        this.style.height = 'auto';
        this.style.height = (this.scrollHeight) + 'px';
    });

})();