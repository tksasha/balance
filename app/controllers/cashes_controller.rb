# frozen_string_literal: true

class CashesController < ApplicationController
  def update
    render :edit, status: :unprocessable_entity if failure?
  end

  private

  def collection
    @collection ||= ::Cashes::GetCollectionService.call(params)
  end

  def result
    @result ||= ::Cashes::GetResultService.call(action_name, params)
  end

  helper_method :result

  delegate :resource, :success?, :failure?, to: :result
end
