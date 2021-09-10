
const clsModalActive = 'modal-active';
const clsModalInActive = 'modal-inactive';

// function to show and hide the modal element
function openModal(modal){
    modal.classList.remove(clsModalInActive);
    modal.classList.add(clsModalActive);
}

function closeModal(modal){
    modal.classList.remove(clsModalActive);
    modal.classList.add(clsModalInActive);
}

function resetControls(){
    const txtRecId = document.getElementById('edit-rec-id');
    const txtAuthor = document.getElementById('edit-value-author');
    const divQuote = document.getElementById('edit-value-quote');
    txtRecId.value = '';
    txtAuthor.value = '';
    while(divQuote.firstChild){
        divQuote.removeChild(divQuote.firstChild);
    }
}

function loadRecord(recId, author, quote){
    const txtRecId = document.getElementById('edit-rec-id');
    const txtAuthor = document.getElementById('edit-value-author');
    const divQuote = document.getElementById('edit-value-quote');
    if (recId != null){
        txtRecId.value = recId;
    }
    if (author != null){
        txtAuthor.value = author;
    }
    if (quote != null){
        divQuote.innerText = quote;
    }
    
}

function initUI(saveAction){
    // get the dom elements
    const btnAddNew = document.getElementById('btn-add-quote');
    const btnEditCancel = document.getElementById('btn-edit-cancel');
    const btnEditSave = document.getElementById('btn-edit-save');
    const divModal = document.getElementById('edit-modal');

    // wire up the event handlers
    btnAddNew.addEventListener('click', ()=> {
        openModal(divModal);
        STATE.setState(ST_VAL_ADD);
        resetControls();
    });
    btnEditCancel.addEventListener('click', ()=> {
        closeModal(divModal);
        resetControls();
        STATE.setState(ST_VAL_READ);
    });
    if (saveAction !== null){
        btnEditSave.addEventListener('click', () => {
           saveAction();
       });
    }
};