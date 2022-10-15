# frozen_string_literal: true

class CashesController < ApplicationController
  private

  def collection
    @collection ||= ::Cashes::GetCollectionService.call(params)
  end

  def result
    @result ||= ::Cashes::GetResultService.call(action_name, params)
  end

  # TODO: spec me
  def dashboard
    @dashboard ||= ::Frontend::Dashboard.new(params)
  end

  # TODO: spec me
  helper_method :dashboard
end
