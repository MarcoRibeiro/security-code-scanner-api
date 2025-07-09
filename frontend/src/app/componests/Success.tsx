"use client";

export const Success = () => {
  return (
    <div className="flex flex-col items-center justify-center p-8">
      <img
        src="https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExd2J0b2J2d2J2b2J2d2J2b2J2d2J2b2J2d2J2b2J2d2J2b2J2/111ebonMs90YLu/giphy.gif"
        alt="No findings!"
        className="w-100 h-100 mb-4"
      />
      <span className="text-white">No findings! Your code is clean</span>
    </div>
  );
};
