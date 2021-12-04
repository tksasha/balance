# frozen_string_literal: true

module ActsAsController
  extend ActiveSupport::Concern

  included do
    helper_method :result

    delegate :resource, :success?, :failure?, to: :result
  end

  def update
    respond_to do |format|
      format.js do
        render :edit, status: :unprocessable_entity if failure?
      end
    end
  end
end
