"use client";

export const Error = () => {
  return (
    <div className="flex flex-col items-center justify-center p-8">
      <img
        src="https://media3.giphy.com/media/v1.Y2lkPTc5MGI3NjExZm5pYjNwMGs3OHNpbGd5d3VpbHhiMzhsd3U3ZWhtdnlpZTl0a285bCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/9PTaAhwri56V2/giphy.gif"
        alt="Error!"
        className="w-100 h-100 mb-4"
      />
      <span className="text-white">Oh no! Something went wrong...</span>
    </div>
  );
};
