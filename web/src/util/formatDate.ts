export function formatDate(date: string | undefined) {
   if (!date) return null;
   return new Date(date).toLocaleDateString();
}