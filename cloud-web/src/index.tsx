import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'

// Extend the Window interface to include isFocused
declare global {
  interface Window {
    isFocused?: boolean;
  }
}

const queryClient = new QueryClient()

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

const handleVisibilityChange = () => window.isFocused = !document.hidden;

const handleWindowBlur = () => {
  window.isFocused = false;
};

const handleWindowFocus = () => {
  window.isFocused = true;
};

document.addEventListener('visibilitychange', handleVisibilityChange);
window.addEventListener('blur', handleWindowBlur);
window.addEventListener('focus', handleWindowFocus);


root.render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <App />
    </QueryClientProvider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
