# frozen_string_literal: true

RSpec.describe TagSerializer do
  subject { described_class.new tag }

  let(:tag) { build :tag, id: 38, name: 'First Tag' }

  its('as_json.symbolize_keys') { should eq id: 38, name: 'First Tag' }
end
