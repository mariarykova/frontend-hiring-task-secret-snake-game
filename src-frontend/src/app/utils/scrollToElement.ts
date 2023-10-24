export const scrollToElement = (id: string) => {
  const element = document.getElementById(id);
  if (element !== null) {
    element.scrollIntoView({ behavior: "smooth" });
  }
};
