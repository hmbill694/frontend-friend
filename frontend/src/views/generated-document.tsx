import NavBar from "../ui/nav/nav";

export type GeneratedDocumentProps = {
	html: string;
};

const foo = { "@click.outside": "$refs.myDialog.close()" };

export default function GeneratedDocument(props: GeneratedDocumentProps) {
	// const [Modal, OpenDialogButton] = Dialog({ dialogId: "chat-drawer", content: "hello" })
	return (
		<>
			<body className="bg-gray-100 min-h-screen flex flex-col" >
				<NavBar />
				<div className="flex flex-col flex-grow" >
					<iframe srcDoc={props.html} className="flex-1"></iframe>
				</div>
			</body>
		</>
	);
}
