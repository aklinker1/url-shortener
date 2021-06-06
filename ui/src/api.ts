import Axios from "axios";

const host = `${window.location.protocol}//${window.location.host}`
const baseUrl = "/@/api";

export interface UrlEntryDto {
  id: string;
  url: string;
  visits: number;
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
    const { data: dtos } = await axios.get<UrlEntryDto[]>("/urlEntries", {
      params: { page, size },
    });
    return dtos.map(this.convertUrlEntryDto);
  },

  async createUrlEntry(url: string): Promise<UrlEntry> {
    const { data: dto } = await axios.post("/urlEntries", { url });
    return this.convertUrlEntryDto(dto);
  },

  async updateUrlEntry(id: string, newUrl: string) {
    const { data: dto } = await axios.put(`/urlEntries/${id}`, {
      url: newUrl,
    });
    return this.convertUrlEntryDto(dto);
  },

  async deleteUrlEntry(id: string) {
    await axios.delete(`/urlEntries/${id}`);
  },

  convertUrlEntryDto(dto: UrlEntryDto): UrlEntry {
    return {
      id: dto.id,
      shortened: `${host}/${dto.id}`,
      url: dto.url,
      visits: dto.visits,
    };
  },
};
