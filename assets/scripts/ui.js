function initUI(saveAction){
    // get the dom elements
    const btnAddNew = document.getElementById('btn-add-quote');
    const divModal = document.getElementById('edit-modal');
    const btnEditCancel = document.getElementById('btn-edit-cancel');
    const btnEditSave = document.getElementById('btn-edit-save');
    const clsModalActive = 'modal-active';
    const clsModalInActive = 'modal-inactive';

    // function to show and hide the modal element
    function openModal(modal){
        if (modal == null) return;
        modal.classList.remove(clsModalInActive);
        modal.classList.add(clsModalActive);
    }

    function closeModal(modal){
        if (modal == null) return;
        modal.classList.remove(clsModalActive);
        modal.classList.add(clsModalInActive);
    }

    function resetControls(){
        const txtAuthor = document.getElementById('edit-value-author');
        const divQuote = document.getElementById('edit-value-quote');
        txtAuthor.value = '';
        while(divQuote.firstChild){
            divQuote.removeChild(divQuote.firstChild);
        }
    }
    
    // wire up the event handlers
    btnAddNew.addEventListener('click', ()=> {
        openModal(divModal);
        STATE.setState(ST_VAL_ADD);
        resetControls();
    });
    btnEditCancel.addEventListener('click', ()=> {
        closeModal(divModal);
        resetControls();
    });
    if (saveAction !== null){
        btnEditSave.addEventListener('click', () => {
           saveAction();
       });
    }
    
};