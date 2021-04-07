import { useQuery } from "@urql/preact";
import { h } from "preact";
import { User } from "./generated/graphql";

const query = `
query GetAllUsers {
  allUsers {
    id
    username
    dateCreated
    messagesCount
  }
}
`;

const Loading = () => {
  return <p>Loading...</p>;
};

const FirstPage = () => {
  const [{ data, fetching }] = useQuery<User[]>({ query });

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
