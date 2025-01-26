import React from "react";
import ReactDOM from "react-dom/client";
import HomePage from "./pages/home/index";
import { BrowserRouter, Route, Routes } from "react-router";
import LoginPage from "./pages/login";
import "./style.css";
import LoginCallbackPage from "./pages/login/callback";

const root = ReactDOM.createRoot(document.getElementById("root"));

root.render(
  <BrowserRouter> 
    <Routes>
      <Route index element={<HomePage />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/login/callback" element={<LoginCallbackPage />} />
    </Routes>
  </BrowserRouter>
);
