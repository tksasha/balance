# frozen_string_literal: true

if Rails.env.development?
  Rails.application.configure do
    config.after_initialize do
      Bullet.enable        = true
      Bullet.alert         = true
      Bullet.bullet_logger = true
      Bullet.console       = true
      Bullet.rails_logger  = true
      Bullet.add_footer    = true
    end
  end
end

if Rails.env.test?
  Rails.application.configure do
    config.after_initialize do
      Bullet.enable        = true
      Bullet.bullet_logger = true
      Bullet.raise         = true
    end
  end
end
