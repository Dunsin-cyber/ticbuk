
## **Project Overview: Ticbuk**

**Ticbuk** is a full-stack application featuring a high-performance **Go** backend, a **PostgreSQL** database for persistent storage, and a cross-platform **React Native** mobile interface.

---

### **Getting Started**

#### **1. Backend (Server)**

The server environment is containerized for consistency. Ensure you have **Docker** installed before proceeding.

* Navigate to the root directory.
* Run the following command to pull dependencies and boot the services:
```bash
make start

```



#### **2. Frontend (Mobile)**

The mobile application is built with Expo.

* Navigate to the `/mobile` directory.
* Install dependencies:
```bash
yarn install

```


* Start the development server:
```bash
npx expo start --tunnel

```



---

### **Connecting Your Device (Network Configuration)**

To test the app on a physical device (iPhone/Android), the mobile client must be able to communicate with your local Go server.

1. **Find your Local IP:** Open your terminal and run `ipconfig`. Look for your **IPv4 Address** (e.g., `192.168.0.x`).
2. **Update API Settings:** Open `mobile/services/api.ts` and update the base URL with your system's IP address:

```typescript
// mobile/services/api.ts
import { Platform } from 'react-native';

const url = Platform.OS === "android" 
    ? "http://10.0.2.2:3000" // Standard Android Emulator loopback
    : "http://192.168.0.X:3000"; // Replace with your actual Local IP address

```

> **Note:** Using your local IP instead of `localhost` allows external physical devices on the same Wi-Fi network to reach your backend API.

---
