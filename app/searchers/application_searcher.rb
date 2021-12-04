# frozen_string_literal: true

class ApplicationSearcher
  attr_reader :params

  def initialize(relation, params = {})
    @relation = relation

    @params = params
  end

  def search
    return results unless params.respond_to? :each

    params.each do |attribute, value|
      method_name = :"search_by_#{ attribute }"

      @results = send method_name, value if respond_to?(method_name, true)
    end

    results
  end

  private

  def results
    @results ||= @relation
  end

  class << self
    def search(*args)
      new(*args).search
    end

    alias call search
  end
end
