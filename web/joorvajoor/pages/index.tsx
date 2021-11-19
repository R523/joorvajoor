import type { NextPage } from "next";
import Head from "next/head";
import styles from "../styles/Home.module.css";

const Home: NextPage = () => {
  const control = (event: string) => {
    return () => {
      fetch(`/api/${event}`)
        .then((resp) => {
          if (!resp.ok) {
            throw new Error(`request failed with ${resp.statusText}`);
          }
        })
        .catch((err) => console.log(err));
    };
  };

  return (
    <div className={styles.container}>
      <Head>
        <title>JoorVaJoor</title>
      </Head>
      <button onClick={control("play")}>Play</button>
      <button onClick={control("pause")}>Pause</button>
      <button onClick={control("volume-up")}>Vol+</button>
      <button onClick={control("volume-down")}>Vol-</button>
    </div>
  );
};

export default Home;
