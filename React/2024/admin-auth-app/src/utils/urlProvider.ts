export interface Provider {
  init(): void;
}

export class UrlProvider implements Provider {
  private _defaultApUrl: string;

  constructor() {
    this._defaultApUrl = "";
    this.init();
  }

  init() {
    this._defaultApUrl = import.meta.env.VITE_API_URL!;
  }

  getBackUrl(): string {
    return this._defaultApUrl;
  }
}

export class UrlProviderInstanceManager {
  private static _instance?: UrlProvider;
  static getInstance(): UrlProvider {
    if (!UrlProviderInstanceManager._instance)
      UrlProviderInstanceManager._instance = new UrlProvider();
    return UrlProviderInstanceManager._instance;
  }
}
// Create an instance and assign it to a unique ad simple variable
export const urlProvider = UrlProviderInstanceManager.getInstance();
