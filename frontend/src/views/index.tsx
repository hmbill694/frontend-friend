import MainLayout from "../layouts/main-layout";

export default function Index() {
  return (
    <>
      <h1 className="text-3xl font-semibold mb-4">Welcome to T0</h1>
      <form action="/api/v1/generate-document" method="post">
        <div className="mb-4">
          <label
            htmlFor="chat"
            className="block text-gray-700 text-sm font-bold mb-2"
          >
            Chat:
          </label>
          <textarea
            id="chat"
            name="chat"
            rows={5}
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          ></textarea>
        </div>
        <div className="flex items-center justify-between">
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit"
          >
            Generate
          </button>
        </div>
      </form>
    </>
  );
}
