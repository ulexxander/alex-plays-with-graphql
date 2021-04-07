import { h } from "preact";
import { useGetAllUsersQuery } from "./generated/graphql";

const Loading = () => {
  return <p>Loading...</p>;
};

const FirstPage = () => {
  const [{ data, fetching }] = useGetAllUsersQuery();

  if (fetching) {
    return <Loading />;
  }

  return (
    <div>
      <h2>Some page</h2>

      <pre>{JSON.stringify(data, null, 4)}</pre>
    </div>
  );
};

export const App = () => {
  return <FirstPage />;
};
