// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyBbzwPZgplM-DvoD_jvsMZeiS2A9WOFmf0",
  authDomain: "test-msg-f5587.firebaseapp.com",
  projectId: "test-msg-f5587",
  storageBucket: "test-msg-f5587.appspot.com",
  messagingSenderId: "750178923326",
  appId: "1:750178923326:web:9d37c514cc3cda3f99726f",
  measurementId: "G-LPPM4F80MD"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);