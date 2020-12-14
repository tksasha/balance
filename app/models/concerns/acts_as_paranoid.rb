# frozen_string_literal: true

module ActsAsParanoid
  extend ActiveSupport::Concern

  included do
    default_scope -> { where(deleted_at: nil) }
  end

  def destroy
    touch(:deleted_at)
  end
end
