import dayjs from 'dayjs';

export function formatDate(date: string | undefined) {
   if (!date) return null;
   return new Date(date).toLocaleDateString();
}

export function parseDateOrNull(date: string | undefined) {
   if (!date) return null;
   return dayjs(date);
}