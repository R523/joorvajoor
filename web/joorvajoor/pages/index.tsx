import React, { useEffect, useState } from "react";
import type { NextPage } from "next";
import Head from "next/head";
import styles from "../styles/Home.module.css";

const Home: NextPage = () => {
  let [isAdmin, setIsAdmin] = useState(false);

  useEffect(() => {
    const timer = setInterval(() => {
      fetch("/api/distance").then((resp) => {
        if (resp.ok) {
          setIsAdmin(true);
        }
      });
    }, 1000);

    return () => clearInterval(timer);
  });

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

      {isAdmin ? <p>You are not admin</p> : <p> you are admin </p>}
    </div>
  );
};

export default Home;
