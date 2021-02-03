import Axios from "axios";

const baseUrl =
  process.env.NODE_ENV !== "development"
    ? "https://apk.rip"
    : "http://localhost:3000";

export interface UrlEntryDto {
  ID: string;
  URL: string;
  Visits: number;
}
export interface UrlEntry {
  id: string;
  shortened: string;
  url: string;
  visits: number;
}

const axios = Axios.create({
  baseURL: baseUrl,
});

export default {
  async listUrlEntries(page: number, size: number): Promise<UrlEntry[]> {
    const { data: dtos } = await axios.get<UrlEntryDto[]>("/api/urlEntries", {
      params: { page, size },
    });
    return dtos.map(this.convertUrlEntryDto);
  },

  async createUrlEntry(url: string): Promise<UrlEntry> {
    const { data: dto } = await axios.post("/api/urlEntries", { url });
    return this.convertUrlEntryDto(dto);
  },

  async updateUrlEntry(id: string, newUrl: string) {
    const { data: dto } = await axios.put(`/api/urlEntries/${id}`, {
      url: newUrl,
    });
    return this.convertUrlEntryDto(dto);
  },

  async deleteUrlEntry(id: string) {
    await axios.delete(`/api/urlEntries/${id}`);
  },

  convertUrlEntryDto(dto: UrlEntryDto): UrlEntry {
    return {
      id: dto.ID,
      shortened: `${baseUrl}/${dto.ID}`,
      url: dto.URL,
      visits: dto.Visits,
    };
  },
};
