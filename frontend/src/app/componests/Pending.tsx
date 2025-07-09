"use client";

export const Pending = () => {
  return (
    <div className="flex flex-col items-center justify-center p-8">
      <img
        src="https://media1.giphy.com/media/v1.Y2lkPTc5MGI3NjExd3Ricjhib2o2OG4zbWVpOXgyM3d0a3Z1bmdyY2ZkeW5lbXh4M2FzdCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/QBd2kLB5qDmysEXre9/giphy.gif"
        alt="Loading..."
        className="w-100 h-100 mb-4"
      />
      <span className="text-white">Scanning in progress...</span>
    </div>
  );
};
