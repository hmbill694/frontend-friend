
export type DialogProps = {
  dialogId: string;
  content: string;
};

function OpenDialogButton(props: DialogProps) {
  return (
    <button
      x-on:click={`$refs.${props.dialogId}.showModal()`}
      id={`${props.dialogId}-open-dialog-button`}
    >
      Open dialog
    </button>
  );
}

export default function Dialog(props: DialogProps) {
  return (
    <div x-data>
      <button
        className="btn btn-primary"
        x-on:click={`$refs.${props.dialogId}.showModal()`}
        id={`${props.dialogId}-open-dialog-button`}
      >
        Open dialog
      </button>
      <dialog x-ref={props.dialogId} className="m-auto card">
        <div className="card-body">
          <h2 className="card-title">Dialog Title</h2>
          <p>{props.content}</p>
          <div className="card-actions justify-end">
            <button
              x-on:click={`$refs.${props.dialogId}.close()`}
              id={`${props.dialogId}-shut-dialog-button`}
            >
              Close
            </button>
          </div>
        </div>
      </dialog>
    </div>
  );
}
