# frozen_string_literal: true

module ActsAsController
  extend ActiveSupport::Concern

  included do
    helper_method :result # TODO: WTF?!

    delegate :resource, :success?, :failure?, to: :result # TODO: WTF?!

    before_action -> { response.status = 201 }, only: :create
  end

  def create
    respond_to do |format|
      format.js do
        render :new, status: :unprocessable_entity if failure?
      end
    end
  end

  def update
    respond_to do |format|
      format.js do
        render :edit, status: :unprocessable_entity if failure?
      end
    end
  end

  def destroy
    respond_to do |format|
      format.js do
        head status: :unprocessable_entity if failure?
      end
    end
  end
end
