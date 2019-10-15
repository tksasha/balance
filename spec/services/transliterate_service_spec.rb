# frozen_string_literal: true

RSpec.describe TransliterateService do
  subject { described_class.transliterate word }

  context do
    let(:word) { 'Алушта' }

    it { should eq 'Alushta' }
  end

  context do
    let(:word) { 'Андрій' }

    it { should eq 'Andrii' }
  end

  context do
    let(:word) { 'Борщагівка' }

    it { should eq 'Borshchahivka' }
  end

  context do
    let(:word) { 'Єнакієве' }

    it { should eq 'Yenakiieve' }
  end

  context do
    let(:word) { 'Короп’є' }

    it { should eq 'Koropie' }
  end

  context do
    let(:word) { 'Згорани' }

    it { should eq 'Zghorany' }
  end

  context do
    let(:word) { 'згорани' }

    it { should eq 'zghorany' }
  end

  context do
    let(:word) { 'Біла Церква' }

    it { should eq 'Bila Tserkva' }
  end

  context do
    let(:word) { nil }

    it { should eq nil }
  end
end
