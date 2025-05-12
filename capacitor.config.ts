import type { CapacitorConfig } from '@capacitor/cli';

const config: CapacitorConfig = {
  appId: 'com.sports.app',
  appName: 'Sports App',
  webDir: 'dist/spa',
  server: {
    androidScheme: 'https',
  },
  plugins: {
    CodePush: {
      androidDeploymentKey: 'YOUR_ANDROID_DEPLOYMENT_KEY',
      iosDeploymentKey: 'YOUR_IOS_DEPLOYMENT_KEY',
      serverUrl: 'https://codepush.appcenter.ms/',
    },
    LiveUpdate: {
      publicKey:
        '-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0vx7agoebGcQSuuPiLJX\nZptN9nndrQmbXEps2aiAFbWhM78LhWx4cbbfAAtVT86zwu1RK7aPFFxuhDR1L6tS\noc/6es03NtnrbtxnMXLxMq6MgUQinhoVILaII6kRSe3J3I7ZqGd4p3yYdJZJxHh\n-----END PUBLIC KEY-----',
      serverDomain: 'https://redamancy.com.cn',
      defaultChannel: 'production',
      autoDeleteBundles: true,
      httpTimeout: 30000,
      readyTimeout: 5000,
      // @ts-expect-error - checkInterval is supported by the plugin but not in the type definition
      checkInterval: 180000, // 3分钟检查一次（单位：毫秒）
    },
  },
};

export default config;
