
export type GeneratedDocumentsProps = {
  files: { projectId: string, id: string, name: string }[];
};

export default function GeneratedDocuments(props: GeneratedDocumentsProps) {
  return (
      <>
        <h1 className="text-3xl font-semibold mb-4">Your Generated Documents</h1>
        <ul>
          {props.files.map((ele) => (
            <li>
              <a href={`/generated-doc/${ele.projectId}/${ele.id}`}>
                {ele.name}
              </a>
            </li>
          ))}
        </ul>
      </>
  );
}
