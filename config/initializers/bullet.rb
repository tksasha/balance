# frozen_string_literal: true

if Rails.env.development?
  Bullet.enable = true

  Bullet.bullet_logger = true

  Bullet.raise = true
end
