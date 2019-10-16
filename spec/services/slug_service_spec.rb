# frozen_string_literal: true

RSpec.describe SlugService do
  subject { described_class.new word }

  context do
    let(:word) { 'Згорани' }

    its(:build) { should eq 'zghorany' }
  end

  context do
    let(:word) { 'Біла Церква' }

    its(:build) { should eq 'bila-tserkva' }
  end

  context do
    let(:word) { '' }

    its(:build) { should eq nil }
  end

  context do
    let(:word) { nil }

    its(:build) { should eq nil }
  end

  describe '.build' do
    subject { described_class.build 'Короп’є' }

    it { should eq 'koropie' }
  end
end
